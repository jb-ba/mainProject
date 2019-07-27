# Installtion

## Cloud
This are all the commands used to install the master in the cloud.

### prepare system
sudo apt update && \
sudo apt install -y snapd curl docker.io && \
sudo systemctl enable docker && \
sudo apt-get install -y openssh-server && \
sudo apt update && \
sudo snap install helm --classic

### Initialize the master and add user to admins
sudo kubeadm init --pod-network-cidr=10.244.0.0/16 >> joinConfigs
mkdir -p $HOME/.kube && \
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config && \
sudo chown $USER $HOME/.kube/config

### Additional components, Flennel and Dashboard
sudo kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
sudo kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/master/aio/deploy/recommended/kubernetes-dashboard.yaml

## Edge (Raspberry Pi)

### Turn swap off && give all cgroups access to cpuset and memory. 
sudo dphys-swapfile swapoff && sudo dphys-swapfile uninstall && sudo update-rc.d dphys-swapfile remove && \
sudo sed -i '$a cgroup_enable=cpuset cgroup_enable=memory' /boot/cmdline.txt

### Install docker/kubeadm dependencies
sudo apt-get install apt-transport-https ca-certificates software-properties-common -y

### Install docker (import keys, setup repo and init)
curl -fsSL get.docker.com -o get-docker.sh && sh get-docker.sh
sudo curl https://download.docker.com/linux/raspbian/gpg
sudo sed -i '$a deb https://download.docker.com/linux/raspbian/ stretch stable' /etc/apt/sources.list
sudo apt-get update
sudo apt-get upgrade -y
sudo systemctl start docker.service

### Install k8s
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt update -q && sudo apt install -y kubeadm kubectl kubelet

### Reboot to load swap and cgroups configs
sudo reboot
