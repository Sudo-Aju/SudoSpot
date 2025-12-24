class Sudospot < Formula
  desc "SudoSpot: Spotify Terminal Player"
  homepage "https://github.com/Sudo-Aju/sudospot"
  url "https://github.com/Sudo-Aju/sudospot.git", branch: "main"
  version "1.0.0"
  head "https://github.com/Sudo-Aju/sudospot.git"

  depends_on "go" => :build

  def install
    system "go", "build", "-o", bin/"sudospot", "cmd/sudospot/main.go"
  end

  test do
    assert_match "SudoSpot", shell_output("#{bin}/sudospot --help", 1)
  end
end
