class GoJames < Formula
    desc "James is your butler and helps you to create, build, debug, test and run your Go projects"
    homepage "https://github.com/pieterclaerhout/go-james"
    url "https://github.com/pieterclaerhout/go-james/releases/download/v1.5.2/go-james_darwin_amd64.tar.gz"
    sha256 "93dd08bf97d1369bd0497c01b2d014511036a728f0b2f9da61f7f3e03a2683de"
    version "1.5.2"
  
    bottle :unneeded
  
    def install
      bin.install "go-james"
    end
  end