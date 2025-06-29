class Freshdocs < Formula
  desc "Keep your documentation as fresh as your code"
  homepage "https://github.com/dullaz/freshdoc"
  version "v1.1.0"
  
  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdoc/releases/download/v1.1.0/freshdocs-v1.1.0-darwin-arm64.tar.gz"
      sha256 "2d5e4e2efcfaaeba6092bfd7cc839839379194cea18d1e31ed91ed6e2a3a7376"
    else
      url "https://github.com/dullaz/freshdoc/releases/download/v1.1.0/freshdocs-v1.1.0-darwin-amd64.tar.gz"
      sha256 "4d4143b1ddd8a05374f9ef4c06635d45badece3145c10e52022cc2d6a610fdae"
    end
  end
  
  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/dullaz/freshdoc/releases/download/v1.1.0/freshdocs-v1.1.0-linux-arm64.tar.gz"
      sha256 "1613a1b52c4e96cb4e099bece4626a06cbae42110c72a484a6a2c8d12eb3b0f9"
    else
      url "https://github.com/dullaz/freshdoc/releases/download/v1.1.0/freshdocs-v1.1.0-linux-amd64.tar.gz"
      sha256 "701b482f4ea6a5411efe655771cd71bda50d5b70cc4db83829c93debdb73b0e5"
    end
  end
  
  def install
    bin.install "freshdocs"
  end
  
  test do
    system "#{bin}/freshdocs", "--version"
  end
end
