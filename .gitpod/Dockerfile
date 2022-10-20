FROM gitpod/workspace-full
# Install KubeBuilder
# @link https://book.kubebuilder.io/quick-start.html
RUN curl -L -o kubebuilder https://go.kubebuilder.io/dl/latest/$(go env GOOS)/$(go env GOARCH) \
  && chmod +x kubebuilder \
  && sudo mv kubebuilder /usr/local/bin/ \
  && sudo sh -c "kubebuilder completion bash > /etc/bash_completion.d/kubebuilder" \
  && kubebuilder version
# Install Kubectl
RUN curl -L "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" -o /tmp/kubectl \
 && sudo install -o root -g root -m 0755 /tmp/kubectl /usr/local/bin/kubectl \
 && sudo sh -c "kubectl completion bash > /etc/bash_completion.d/kubectl" \
 && kubectl version --client
