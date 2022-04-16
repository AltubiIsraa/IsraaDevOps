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
							archiveArtifacts artifacts: 'sample', followSymlinks: false   
                     }
                }

        }
}
