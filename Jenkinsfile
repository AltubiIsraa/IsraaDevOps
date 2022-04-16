
pipeline{
        agent any
                stages{
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

        }
}
