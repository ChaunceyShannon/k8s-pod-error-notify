package main

import (
	"bytes"
	"context"
	"io"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type argStruct struct {
	InCluster  bool
	ConfigFile string
	Namespace  string
	Webhook    string
}

func main() {
	args := new(argStruct)
	a := argparser("kubernetes的pod程序崩溃通知程序")
	args.InCluster = a.getBool("", "InCluster", "false", "是否在集群内部, 如果不在集群内部需要指定config文件")
	args.ConfigFile = a.get("", "ConfigFile", "", "如果不在集群内部, 需要指定配置文件")
	args.Namespace = a.get("", "Namespace", "", "需要监听事件的namespace, 逗号分割, 默认为空, 即监听所有")
	args.Webhook = a.get("", "Webhook", "", "webhook的url, 会post到这个url, body就是消息内容")
	a.parseArgs()

	if !args.InCluster {
		if args.ConfigFile == "" {
			lg.error("在集群外部需要指定配置文件")
			exit(1)
		}
		if !pathExists(args.ConfigFile) {
			lg.error("配置文件不存在")
			exit(1)
		}
	}

	lg.setLevel("debug")

	lg.info("开始")

	var nss []string
	if args.Namespace != "" {
		nss = strSplit(args.Namespace, ",")
	} else {
		nss = append(nss, "")
	}

	for _, ns := range nss {
		go func(ns string) {
			if ns == "" {
				lg.trace("监听命名空间: 所有")
			} else {
				lg.trace("监听命名空间:", ns)
			}
			for {
				var config *rest.Config
				var err error
				if args.InCluster {
					lg.trace("在集群内部")
					config, err = rest.InClusterConfig()
					panicerr(err)
				} else {
					lg.trace("在集群外部")
					config, err = clientcmd.BuildConfigFromFlags("", args.ConfigFile)
					panicerr(err)
				}

				client, err := kubernetes.NewForConfig(config)
				panicerr(err)

				watcher, err := client.CoreV1().Pods(ns).Watch(context.Background(), metav1.ListOptions{})
				panicerr(err)

			loopier:
				for {
					select {
					case event, ok := <-watcher.ResultChan():
						if !ok {
							lg.trace("事件event关闭, 重新开始监听")
							sleep(1)
							break loopier
						}
						try(func() {
							pod := event.Object.(*v1.Pod)
							lg.trace("监听到事件:", event.Type, pod.ObjectMeta.Namespace, pod.ObjectMeta.Name)

							for _, cs := range pod.Status.ContainerStatuses {
								name := cs.Name
								if cs.State.Terminated != nil && cs.State.Terminated.Reason == "Error" {
									rs, err := client.CoreV1().Pods(pod.ObjectMeta.Namespace).GetLogs(pod.ObjectMeta.Name, &v1.PodLogOptions{
										Container: name,
										Follow:    false,
										TailLines: func(a int64) *int64 {
											return &a
										}(50),
									}).Stream(context.Background())

									panicerr(err)

									buf := new(bytes.Buffer)
									_, err = io.Copy(buf, rs)
									panicerr(err)
									s := buf.String()

									res := ""
									for _, ss := range strSplit(s, "\n") {
										if strStartsWith(ss, "panic: ") {
											res = ss
										} else if res != "" {
											res = res + "\n" + ss
										}
									}

									msg := "Namespace: " + pod.ObjectMeta.Namespace + "\nPod: " + pod.ObjectMeta.Name + "\nContainer: " + name + "\n"
									if res != "" {
										msg = msg + "Golang Panic Stacktrace: \n\n"
										msg = msg + substr(res, len(res)-3000, len(res))
									} else {
										msg = msg + "Lastest Log:\n\n"
										msg = msg + substr(s, len(s)-3000, len(s))
									}

									lg.info(msg)
									if err := try(func() {
										httpPostRaw(args.Webhook, msg)
									}, tryConfig{retry: 3, sleep: 3}).Error; err != nil {
										lg.error(err)
									}
								}
							}
						}).except(func(err error) {
							lg.warn(err)
						})
					default:
						sleep(0.1)
					}
				}
			}
		}(ns)
	}
	select {}
}
