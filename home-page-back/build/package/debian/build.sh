BASE_DIR="/tmp/home-page-back"
DEB_FILE="/tmp/home-page-back.deb"


currentDit=$(pwd)
cd /home-page-back/cmd/home-page
rm home-page
go build
if [[ ! -x home-page ]]; then
    echo "go buil failed"
    exit 1
fi

if [[ -d "${BASE_DIR}" ]]; then
    rm -rfv "${BASE_DIR}"
fi

mkdir -pv "${BASE_DIR}/DEBIAN"
cp -v /home-page-back/build/package/debian/control "${BASE_DIR}/DEBIAN"

mkdir -pv "${BASE_DIR}/lib/systemd/system"
cp -v /home-page-back/build/package/debian/home-page.service "${BASE_DIR}/lib/systemd/system"

mkdir -pv "${BASE_DIR}/etc/home-page-back"
mkdir -pv "${BASE_DIR}/var/log/home-page-back"
mkdir -pv "${BASE_DIR}/opt/ltt/home-page-back/dbs"
cp -v home-page "${BASE_DIR}/opt/ltt/home-page-back"

dpkg-deb --build "${BASE_DIR}"
rm -rfv "${BASE_DIR}"
echo "Debian file builded: '${DEB_FILE}'"
cd ${currentDit}