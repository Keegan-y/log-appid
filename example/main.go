package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//一起从这开始
	ns,err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	//fmt.Printf("lenth:%v", len(ns.Items))
	for i:=0;i<len(ns.Items);i++ {
		//fmt.Println(ns.Items[i].Name)
		if strings.HasPrefix(ns.Items[i].Name, "ns") {
			//bool
			namespace := ns.Items[i].Name
			pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}
			//fmt.Printf("Pod  in namespace %s not found\n", namespace)
			//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
			for j:=0;j<len(pods.Items);j++{
				pod := pods.Items[j].Name
				//fmt.Printf("%v\n",pod[0:40])
				if strings.HasPrefix(pod, "app"){
					fmt.Printf("%v\n",pod[0:40])
				} else {
					fmt.Printf("%v",pod)
				}

				//fmt.Println(pod)//打印pod id总长度
				//cls-a91a7dbc-1a23-42f6-a66f-876094349f03
				//模仿ns.Items[i].Name
				/*3 8 4 4 4 6 5
				_, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
				if errors.IsNotFound(err) {
					fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
				} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
					fmt.Printf("Error getting pod %s in namespace %s: %v\n",
						pod, namespace, statusError.ErrStatus.Message)
				} else if err != nil {
					panic(err.Error())
				} else {
					fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
				}
				time.Sleep(3 * time.Second)

				 */
			}

		}
	}
}


func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}