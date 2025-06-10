pipeline{
    agent any 
    stages{
        stage('Checkout') {
            steps{
                git 'https://github'
            }
        }

        stage('Test'){
            steps{
                sh 'go test -v'
            }
        }

        stage('Build'){
            steps{
                sh 'go build -o myapp'
            }
        }

        stage('Run'){
            steps{
                sh './myapp'
            }
        }
    }
}