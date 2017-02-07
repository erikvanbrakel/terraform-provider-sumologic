require "rake/clean"

CLEAN.include 'pkg','dist'

directory "dist"

task :default => [:fmtcheck]

fmt_files = Dir.glob("./**/*.go").join " "

desc "Verify code passes go fmt"
task :fmtcheck do
    fmt_result = `gofmt -l #{fmt_files}`
    if (fmt_result.length > 0) then
        puts 'gofmt needs running on the following files:'
        fmt_result.split("\n").each do |f|
            puts "\t#{f}"
        end
        puts ""
        puts "You can use the command: \`rake fmt\` to reformat code."
        exit 1
    else
        puts "All files are formatted correctly!"
    end
end

task :fmt do
    sh "gofmt -w #{fmt_files}"
end

namespace "build" do
    target_os = "linux darwin windows freebsd openbsd solaris"
    target_arch = "amd64 386 arm"
    exclude_os_arch = "!darwin/arm !darwin/386"

    task :tools do
        sh "go get -u github.com/mitchellh/gox"
    end

    desc "Build binaries for all platforms"
    task :all => [:tools, :clean] do
        sh "gox -os=\"#{target_os}\" -verbose -arch=\"#{target_arch}\" -osarch=\"#{exclude_os_arch}\" -output \"./pkg/{{.OS}}_{{.Arch}}/terraform-provider-sumologic\" ."
    end

    desc "Build binary for current platform only"
    task :dev => [:tools, :clean] do
    end

    desc "Package binaries into zip files"
    task :package => [:all, "dist"] do
        Dir.chdir("pkg") do
            dirs = Dir["*"].reject{|o| not File.directory?(o)}
            dirs.each do |dir|
                Dir.chdir(dir) do
                    sh "zip ../../dist/#{dir}.zip *"
                end
            end
        end
    end
end
