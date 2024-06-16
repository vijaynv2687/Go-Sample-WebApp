pipeline {
  // Run on an agent where we want to use Go
  agent any

  // Ensure the desired Go version is installed for all stages,
  // using the name defined in the Global Tool Configuration
  tools { 
      go '1.22.0' 
      
  }

    environment {
        GO111MODULE='on'
    }

  stages {
    stage('Build') {
      steps {
          git branch: 'main', 
                    url: 'https://github.com/vijaynv2687/Go-Sample-WebApp.git'
       sh 'go build .'
      }
    }
    stage('RUN'){
        steps {
            sh 'cd /Users/vijaynv/.jenkins/workspace/goapp && nohup go run main.go'
        }
    }
  }
}
