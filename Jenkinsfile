pipeline{
        agent any
            stages{
              stage('Test'){
                  steps{
                      sh 'go test'
                     }
                }
               stage('Build'){
                   steps{
                      sh 'go build -o sample1 main.go'
                        }
                 }
                stage('Save srtifact'){
                    steps{
                      archiveArtifacts artifacts: 'sample1', followSymlinks: false
                     }
                 }
	          stage('Docker login'){
                 steps {
			 withCredentials([usernamePassword(credentialsId: 'devops-docker-creds', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
               		  sh 'docker login --username ${USERNAME} --password ${PASSWORD}'
                      }
		    }
                 }
	          stage('Docker build'){
                 steps{
                  sh 'docker build . --tag altubiisraa97/devops:v2'
                      }		     
                 }
	          stage('Docker push'){
                 steps{
		     sh 'docker push altubiisraa97/devops:v2'
                      }
                 }
        }
}
