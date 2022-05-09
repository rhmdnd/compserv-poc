# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "fedora/35-cloud-base"
  config.vm.synced_folder ".", "/vagrant", type: "virtualbox"
  config.vm.provider "virtualbox" do |vb|
    vb.memory = "4096"
  end
end
