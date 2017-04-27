Vagrant.configure("2") do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.provision :ansible do |ansible|
    ansible.playbook = "vagrant.yml"
    ansible.verbose = true
  end

  config.vm.network "private_network", type:"dhcp"
end
