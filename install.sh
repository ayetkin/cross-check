remove(){
  echo -e "Stopping service..." && sleep 2 & spinner
  systemctl stop cross-check.service
  echo -e "Reloading system daemon..." && sleep 2 & spinner
  systemctl daemon-reload
  echo -e "Deleting configs..." && sleep 2 & spinner
  rm -rf /etc/cross-check
  echo -e "Deleting system service..." && sleep 2 & spinner
  rm -rf /etc/systemd/system/cross-check.service
  echo -e "Deleting cross check binary..." && sleep 2 & spinner
  rm -rf /usr/local/bin/cross-check
  echo -e ":::::::: Installing done ::::::::"
}

update () {
  if test -f "/usr/local/bin/cross-check"; then
    echo -e "Another cross check version exists. Switching upgrade mode..." && sleep 2 & spinner
    # shellcheck disable=SC2216
    echo -e "Backup installed version to tmp directory..." && sleep 2 & spinner
    # shellcheck disable=SC2216
    yes | cp -r /usr/local/bin/cross-check /tmp/cross-check-backup
    echo -e "Stopping service..." && sleep 2 & spinner
    systemctl stop cross-check.service
    echo -e "Copying binary..." && sleep 2 & spinner
    # shellcheck disable=SC2216
    yes | cp -r release/cross-check /usr/local/bin/cross-check
    echo -e "Starting service..." && sleep 2 & spinner
    systemctl start cross-check
    if [ $? -ne 0 ]; then
      echo "Error while stopping service"
      exit 1
    fi
    echo -e ":::::::: Upgrading done ::::::::"
  else
    echo -e "Can not detected existsing Cross Check version!" && exit 1
  fi
}

install () {
  echo -e "Creating config directory (/etc/cross-check)..." && sleep 2 & spinner
  mkdir -p /etc/cross-check

  echo -e "Copying binary..." && sleep 2 & spinner
  # shellcheck disable=SC2216
  yes | cp -r release/cross-check /usr/local/bin/cross-check

  if [ $1 == "mysql" ]; then
    echo -e "Copying MYSQL service..." && sleep 2 & spinner
    # shellcheck disable=SC2216
    yes | cp -r release/cross-check-mysql.service /etc/systemd/system/cross-check.service
  elif [ $1 == "pgsql" ]; then
    echo -e "Copying PGSQL service..." && sleep 2 & spinner
    # shellcheck disable=SC2216
    yes | cp -r release/cross-check.service /etc/systemd/system/cross-check.service
  fi

  echo -e "Copying config file..." && sleep 2 & spinner
  # shellcheck disable=SC2216
  yes | cp -r release/config.yaml /etc/cross-check/

  echo -e "Reloading system daemon..." && sleep 2 & spinner
  systemctl daemon-reload
  echo -e "Enabling service..." && sleep 2 & spinner
  systemctl enable cross-check
  echo -e "Starting service..." && sleep 2 & spinner
  systemctl start cross-check
  # shellcheck disable=SC2181
  if [ $? -ne 0 ]; then
      echo "Error while starting service"
      exit 1
  fi
  echo -e ":::::::: Installing done ::::::::"
}

spinner()
{
    local pid=$!
    local delay=0.75
    local spinstr='|/-\'
    while [ "$(ps a | awk '{print $1}' | grep $pid)" ]; do
        local temp=${spinstr#?}
        printf " [%c]  " "$spinstr"
        local spinstr=$temp${spinstr%"$temp"}
        sleep $delay
        printf "\b\b\b\b\b\b"
    done
    printf "    \b\b\b\b"
}

#Copy service config and application

if [ $1 == "mysql" ]; then
  echo -e ":::::::: Cross Check MYSQL Installing started ::::::::" && sleep 1
  install $1
elif [ $1 == "pgsql" ]; then
  echo -e ":::::::: Cross Check PGSQL Installing started ::::::::" && sleep 1
  install
elif [ $1 == "update" ]; then
  echo -e ":::::::: Cross Check Upgrading started ::::::::" && sleep 1
  update
elif [ $1 == "remove" ]; then
  echo -e ":::::::: Cross Check Uninstalling ::::::::" && sleep 1
  remove
else
  echo -e "Please select an argument. (mysql, pgsql, update, delete)" && exit 1
fi

