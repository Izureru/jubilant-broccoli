task :build => ["get"]do 
	sh "go build"
end

task :get do 
	sh "go get"
end

namespace :run do
	
	task :local => ["build", "config:local"] do
		sh  "./vitamns"
	end

	task :dev => ["build", "config:dev"] do
		sh  "./vitamns"
	end
end

namespace :config do

	task :local do
		ENV['PORT']="1212"
		ENV['MONGOURI']="http://localhost:1212"
		ENV['DBNAME']="simplyactivedb"
		ENV['CNAME']="mnshealth"
		ENV['APIKEY']="1234"
	end

	task :mimic do
		ENV['PORT']="3800"
		ENV['MONGOURI']="127.0.0.1:27017"
		ENV['DBNAME']="mongoid"
		ENV['CNAME']="meals"
		ENV['APIKEY']="1234"
	end
end