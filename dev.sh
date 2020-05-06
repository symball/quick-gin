#!/bin/bash
#!/
#!/# Author Simon Ball <open-source@simonball.me>
# A short shell script used to bootstrap services

#
# DEFAULTS. Should be no need to change these
#
red=`tput setaf 1`
green=`tput setaf 2`
blue=`tput setaf 4`
reset=`tput sgr0`
FOREGROUND=false
GROUP_ID=false
LAST_RUN_FILE=./.LAST_RUN
LOGS_ATTACH=false
PERFORM_SETUP=false
RUN_FIXTURES=false
RUNNING_LOCK=./.RUNNING
USER_ID=false
VERBOSE=true

#
# ENVIRONMENT SPECIFIC. This is the section you will want to change
#
REQUIRED_PROGRAMS="docker docker-compose"
REQUIRED_PORTS="5432 8080"

#
# PRE SETUP
#
# If going to do something before starting up container / main env
# at setup time
#
pre_setup()
{
  echo "${green}--=== PRE SETUP ===--${reset}"
  docker-compose down -v
  [ -f $RUNNING_LOCK ]; rm $RUNNING_LOCK

}
#
# SETUP
#
# Routine to perform once support services and env file managed
#
setup()
{
  echo "${green}--=== SETUP ===--${reset}"
  
  ./config/postgres_wait.sh
  ok

  POSTGRESQL_URL=postgres://development:development@localhost:5432/development?sslmode=disable
  ./migrate -database ${POSTGRESQL_URL} -path db/migrations up
}
#
# DOWN
#
# Shutting down all services
#
down()
{
  echo "${green}--=== DOWN ===--${reset}"
  [ -f ./docker-compose.yml ] && docker-compose down
  rm $RUNNING_LOCK > /dev/null 2>&1
  halt "User stopped"
}
#
# END CUSTOMISATION
#

logo() {
  cat << "EOF"
 __                 _           _ _      
/ _\_   _ _ __ ___ | |__   ___ | (_) ___ 
\ \| | | | '_ ` _ \| '_ \ / _ \| | |/ __|
_\ \ |_| | | | | | | |_) | (_) | | | (__ 
\__/\__, |_| |_| |_|_.__/ \___/|_|_|\___|
    |___/ 

EOF
}

logo
echo "${green}--=== Symbolic Quick Script ===--${reset}"
echo ""

#
# HALT
#
# Safe method for stopping dev environment. If running the main
# process in foreground, will also shutdown services
#
halt() {
  echo "${red}--=== HALTING ===--${reset}"
  echo "$1"    
  exit 1
}

trap halt SIGHUP SIGINT SIGTERM

# Function to check whether command exists or not
exists()
{
  if command -v $1 &>/dev/null
    then return 0
    else return 1
  fi
}

ok() {
  echo -e " ${green}OK${reset}"
}

# Command help
display_usage() {
  echo "Get a basic environemnt up and running on a local device in a common format"
  echo ""
  echo " -d --database        Run fixtures procedure"
  echo " -h --help            Display this message and exit"
  echo " -s --initial-setup   Run a series off initiation procedures on the project"
  echo " -q --quiet           Minimal output"
  echo " -x --stop            Run the halt procedure"
  halt
}

# Parameter parsing
while [ $# -gt 0 ]; do
  case "$1" in
    --help|-h)
      display_help
      ;;
    --database|-d)
      RUN_FIXTURES=true
      ;;
    --quiet|q)
      VERBOSE=false
      ;;
    --setup|-s)
      PERFORM_SETUP=true
      ;;
    --stop|-x)
      down
      ;;
  esac
  shift
done

#
# BASIC PREPARATION
#

# Check whether the required programs installed
[ "$VERBOSE" = true ] && echo "--- Checking required programs"
for PROGRAM in $REQUIRED_PROGRAMS; do
  if exists $PROGRAM; then
    [ "$VERBOSE" = true ] && echo -ne "$PROGRAM" && ok
  else halt "$PROGRAM Required"
  fi
done

# If the script has never been run before, flip the initial setup condition
if [ ! -f $LAST_RUN_FILE ]; then
  [ "$VERBOSE" = true ] && echo "${green}First run detected${reset}"
  PERFORM_SETUP=true
fi

# Check whether ports are available
if [ ! -f $RUNNING_LOCK ]; then
  [ "$VERBOSE" = true ] && echo "--- Open Ports"
  for PORT in $REQUIRED_PORTS; do
    PORT_RESULT="$(lsof -i :${PORT})"
    if [ -z "$PORT_RESULT"]; then
      [ "$VERBOSE" = true ] && echo -ne "$PORT" && ok
    else
      halt "Port $PORT already in use"
    fi
  done
fi

#
# SERVICE STARTUP
#
echo "OK" > $LAST_RUN_FILE
if $PERFORM_SETUP ; then
  [ "$VERBOSE" = true ] && echo "--- Running Setup"
  if [ -f ./.env.dist ]; then
    echo "Copying environment file"
    cp .env.dist .env
  else 
   echo "" > .env
  fi
  pre_setup
  if [ ! -f $RUNNING_LOCK ]; then
    if [ -f ./docker-compose.yml ]; then
      [ "$VERBOSE" = true ] && echo "Starting Docker"
      docker-compose up -d
      echo "OK" > $RUNNING_LOCK
    fi
  fi
  setup
  ok
else
  # Support services
  if [ ! -f $RUNNING_LOCK ]; then
    if [ -f ./docker-compose.yml ]; then
      [ "$VERBOSE" = true ] && echo "--- Running Docker support services"
      if [ "$FOREGROUND" = true ]; then
        docker-compose up
        echo "OK" > $RUNNING_LOCK
      else
        docker-compose up -d
        echo "OK" > $RUNNING_LOCK
      fi
    fi
  fi
  ./config/postgres_wait.sh
fi
./quick-gin