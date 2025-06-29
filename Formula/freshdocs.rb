class Freshdocs < Formula
  desc "Keep your documentation as fresh as your code"
  homepage "https://github.com/dullaz/freshdocs"
  version "v1.2.0"
  
  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdocs/releases/download/v1.2.0/freshdocs-v1.2.0-darwin-arm64.tar.gz"
      sha256 "bef4bfd96ea02a9f3fc7eb8946725ca7ddb6999a28da68fe89943ecdd0de805e"
    else
      url "https://github.com/dullaz/freshdocs/releases/download/v1.2.0/freshdocs-v1.2.0-darwin-amd64.tar.gz"
      sha256 "ee2eab6b5785080b1fcc45b2ee24cc339eead4c64082d2538de7246e6193c44c"
    end
  end
  
  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdocs/releases/download/v1.2.0/freshdocs-v1.2.0-linux-arm64.tar.gz"
      sha256 "99a9a3a2be84b95faf3ed35f89457e7d05281ed4c92bfd3a59483ee037ae9582"
    else
      url "https://github.com/dullaz/freshdocs/releases/download/v1.2.0/freshdocs-v1.2.0-linux-amd64.tar.gz"
      sha256 "b229154702f0422bdc6aa3645d76007c28c60854ce46e6a7f2a2df665046c512"
    end
  end
  
  def install
    bin.install "freshdocs"
  end
  
  test do
    system "#{bin}/freshdocs", "--version"
  end
end
