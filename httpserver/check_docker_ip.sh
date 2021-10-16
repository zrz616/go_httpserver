HTTPSERVER_PID=`pidof httpserver`
nsenter -t $HTTPSERVER_PID -n ip a
