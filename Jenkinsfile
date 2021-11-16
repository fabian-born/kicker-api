podTemplate(label: 'mypod', containers: [
    containerTemplate(name: 'docker', image: 'docker', ttyEnabled: true, command: 'cat'),
    containerTemplate(name: 'kubectl', image: 'lachlanevenson/k8s-kubectl:v1.22.0', command: 'cat', ttyEnabled: true),
  ],
  volumes: [
    hostPathVolume(mountPath: '/var/run/docker.sock', hostPath: '/var/run/docker.sock'),
  ]) {
    node('mypod') {
        
        deleteDir()
        
        stage("Checkout") {
            checkout scm
        }
        stage('Build and Push Container') {
            container('docker') {

                withCredentials([[$class: 'UsernamePasswordMultiBinding',
                        credentialsId: 'gitlab-credential-fb',
                        usernameVariable: 'DOCKER_HUB_USER',
                        passwordVariable: 'DOCKER_HUB_PASSWORD']]) {
                    
                    sh "docker build -t fabianborn/webapp:${env.BUILD_NUMBER} ."
//                    sh "docker login -u fabianborn -p ${env.DOCKER_HUB_PASSWORD} "
//                   sh "docker push fabianborn/webapp:${env.BUILD_NUMBER} "
                }
            }
        }
        
        stage ("Automated Test Cases"){
          
          // give the container 10 seconds to initialize the web server
          sh "sleep 10"

          // connect to the webapp and verify it listens and is connected to the db
          //
          // to get IP of jenkins host (which must be the same container host where dev instance runs)
          // we passed it as an environment variable when starting Jenkins.  Very fragile but there is
          // no other easy way without introducing service discovery of some sort
          echo "Check if webapp port is listening and connected with db"
          // sh "curl http://192.168.42.5:30001/v1/ping -o curl.out"
          // sh "cat curl.out"
          // sh "awk \'/true/{f=1} END{exit!f}\' curl.out"
          echo "<<<<<<<<<< Access this test build at http://192.168.42.5:30001 >>>>>>>>>>"        
        }
        def push = ""
        stage ("Manual Test & Approve Push to Production"){
          // Test instance is online.  Ask for approval to push to production.
          // notifyBuild('APPROVAL-REQUIRED')
          push = input(
            id: 'push', message: 'Push to production?', parameters: [
              [$class: 'ChoiceParameterDefinition', choices: 'Yes\nNo', description: '', name: 'Select yes or no']
            ]
          )
        }
        
        stage('Deploy in Production') {
            container('kubectl') {
                sh "sleep 10"
                withCredentials([[$class: 'UsernamePasswordMultiBinding', 
                        credentialsId: 'gitlab-credential-fb',
                        usernameVariable: 'DOCKER_HUB_USER',
                        passwordVariable: 'DOCKER_HUB_PASSWORD']]) {
                    
                }
            }
        }
    }

