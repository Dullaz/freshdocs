class Freshdocs < Formula
  desc "Keep your documentation as fresh as your code"
  homepage "https://github.com/dullaz/freshdocs"
  version "v1.3.0"
  
  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdocs/releases/download/v1.3.0/freshdocs-v1.3.0-darwin-arm64.tar.gz"
      sha256 "2ddb8c094cf9641163e0886a961fabee6ccf12c8aa9cbc2b82251b629040aba8"
    else
      url "https://github.com/dullaz/freshdocs/releases/download/v1.3.0/freshdocs-v1.3.0-darwin-amd64.tar.gz"
      sha256 "50f504287c74df947ac03b1c6d7d6c28a33ddbece7562320e9c95e5dec602678"
    end
  end
  
  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdocs/releases/download/v1.3.0/freshdocs-v1.3.0-linux-arm64.tar.gz"
      sha256 "5c7ffa9e9a2ef1d1da71a88b71ee6864f20d84af219227865a1d0133f9034804"
    else
      url "https://github.com/dullaz/freshdocs/releases/download/v1.3.0/freshdocs-v1.3.0-linux-amd64.tar.gz"
      sha256 "0011afe4b3edfcbfe96f12b2c3e1cd89dee95b641e13b4fdd1834c237f6da709"
    end
  end
  
  def install
    bin.install "freshdocs"
  end
  
  test do
    system "#{bin}/freshdocs", "--version"
  end
end
