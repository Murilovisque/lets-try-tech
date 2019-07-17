#!/bin/bash

start_home_page_back() {
	local proc=$(ps a | grep home-page-back | grep -v " grep " | grep -v "/etc/init.d/home-page-back")
	if [[ -z $proc ]]; then
		echo "Starting..."
		/opt/ltt/home-page-back/home-page-back &
		echo "Started"
	else
		echo "Already running..."
	fi
}

stop_home_page_back() {
	ps a | grep home-page-back | grep -v " grep " | awk '{ print $1 }' | head -n 1 | xargs kill
}

status_home_page_back() {
	local proc=$(ps a | grep home-page-back | grep -v " grep " | grep -v "/etc/init.d/home-page-back")
	if [[ -z $proc ]]; then
		echo "home-page-back stopped"
	else
		echo "home-page-back running..."
	fi
}

case "$1" in
	status)
		status_home_page_back
		exit $?
		;;
	start)
		start_home_page_back
		exit $?
		;;
    stop)
		stop_home_page_back
		exit $?
		;;
	restart)
		start_home_page_back
		exit $?
		;;
	*)
		echo "Usage {start|stop|restart}"
		exit 3
		;;
esac