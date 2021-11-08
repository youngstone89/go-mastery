package executable

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func AdminGetConfigNotifyWebhook() {
	command := exec.Command("./mc", "admin", "config", "get", "minio-local/", "notify_webhook")

	// command := exec.Command("./mc", "ls", "minio-local")
	command.Dir = "/usr/local/bin/"
	// fmt.Println(command.Dir)
	// command.Env = []string{"MC_HOST_minio=http://20ee7722-444a-4ff4-afc2-a242bcad1a2b:d10ddd48-3136-4045-8292-c8ee3543957e@localhost:9000"}

	// export MC_HOST_<alias>=https://<Access Key>:<Secret Key>@<YOUR-S3-ENDPOINT>
	// "localhost:9000", "20ee7722-444a-4ff4-afc2-a242bcad1a2b", "d10ddd48-3136-4045-8292-c8ee3543957e"
	// mc admin config get ps-lab-minio/ notify_webhook
	// set var to get the output
	var out bytes.Buffer
	var errb bytes.Buffer

	// set the output to our variable
	command.Stdout = &out
	command.Stderr = &errb
	err := command.Run()
	if err != nil {
		log.Println(errb)
	}

	fmt.Println(out.String())
	fmt.Println(errb.String())
}

func RestartMinioService() {
	command := exec.Command("./mc", "admin", "service", "restart", "minio-local/", "--debug")
	// command := exec.Command("./mc", "ls", "minio-local")
	command.Dir = "/usr/local/bin/"
	// fmt.Println(command.Dir)
	// command.Env = []string{"MC_HOST_minio=http://20ee7722-444a-4ff4-afc2-a242bcad1a2b:d10ddd48-3136-4045-8292-c8ee3543957e@localhost:9000"}

	// export MC_HOST_<alias>=https://<Access Key>:<Secret Key>@<YOUR-S3-ENDPOINT>
	// "localhost:9000", "20ee7722-444a-4ff4-afc2-a242bcad1a2b", "d10ddd48-3136-4045-8292-c8ee3543957e"
	// mc admin config get ps-lab-minio/ notify_webhook
	// set var to get the output
	var out bytes.Buffer
	var errb bytes.Buffer

	// set the output to our variable
	command.Stdout = &out
	command.Stderr = &errb

	go command.Run()

	time.Sleep(5 * time.Second)

	fmt.Println("restarted")

	fmt.Println(out.String())
	fmt.Println(errb.String())
}

func ListBuckets() {
	// command := exec.Command("./mc", "admin", "config", "get", "minio-local/", "notify_webhook")
	command := exec.Command("./mc", "ls", "minio-local")
	command.Dir = "/usr/local/bin/"
	// fmt.Println(command.Dir)
	// command.Env = []string{"MC_HOST_minio=http://20ee7722-444a-4ff4-afc2-a242bcad1a2b:d10ddd48-3136-4045-8292-c8ee3543957e@localhost:9000"}

	// export MC_HOST_<alias>=https://<Access Key>:<Secret Key>@<YOUR-S3-ENDPOINT>
	// "localhost:9000", "20ee7722-444a-4ff4-afc2-a242bcad1a2b", "d10ddd48-3136-4045-8292-c8ee3543957e"
	// mc admin config get ps-lab-minio/ notify_webhook
	// set var to get the output
	var out bytes.Buffer
	var errb bytes.Buffer

	// set the output to our variable
	command.Stdout = &out
	command.Stderr = &errb
	err := command.Run()
	if err != nil {
		log.Println(errb)
	}

	fmt.Println(out.String())
	fmt.Println(errb.String())
}

// mc admin config set minio-local notify_webhook:1 queue_limit="0" endpoint="http://77f5-125-176-145-160.ngrok.io/api/v1/minio/events" queue_dir=""
func SetNotifyWebhook(identifer string, queueLimit string, endpoint string, queue_dir string) {
	command := exec.Command("./mc", "admin", "config", "set", "minio-local",
		fmt.Sprintf("%s:%s", "notify_webhook", identifer),
		fmt.Sprintf("queue_limit=\"%s\"", queueLimit),
		fmt.Sprintf("endpoint=\"%s\"", endpoint),
		fmt.Sprintf("queue_dir=\"%s\"", queue_dir),
	)
	command.Dir = "/usr/local/bin/"

	var out bytes.Buffer
	var errb bytes.Buffer

	// set the output to our variable
	command.Stdout = &out
	command.Stderr = &errb
	err := command.Run()
	if err != nil {
		log.Println(errb)
	}

	fmt.Println(out.String())
	fmt.Println(errb.String())
}
