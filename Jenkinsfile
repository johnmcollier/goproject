#!groovy​

pipeline {
    agent {
    
    kubernetes {        
      label 'codewind-agent-pod'
      yaml """
      apiVersion: v1
      kind: Pod
      spec:
      containers:
          - name: buildah
            image: johncollier/buildah
            command:
            - cat
            tty: true
            securityContext:
              privileged: true
            resources:
              limits:
                memory: "2Gi"
                cpu: "1"
              requests:
                memory: "2Gi"
                cpu: "1"
        """
        }
    }
    
    options {
        timestamps() 
        skipStagesAfterUnstable()
    }
    
    stages {

        stage('Build') {
            container('buildah') {
                steps {
                    script {
                        println("Starting codewind-eclipse build...")
                        
                        def sys_info = sh(script: "uname -a", returnStdout: true).trim()
                        
                        println("System information: ${sys_info}")
                        
                        sh '''
                        buildah bud -t test_build .
                        '''
                    }
                }
            }

        }        
    }    
}
