## data preparation
if you have installed libosrm, you may use it locally, otherwise you may also user the docker
```bash
docker run -t -v $(pwd):/data osrm/osrm-backend osrm-extract -p /opt/car.lua /data/china-latest.osm.pbf
docker run -t -v $(pwd):/data osrm/osrm-backend osrm-partition /data/china-latest.osrm
docker run -t -v $(pwd):/data osrm/osrm-backend osrm-customize /data/china-latest.osrm
docker-compose run osrm/osrm-backend osrm-routed --algorithm mld /data/china-latest.osrm
```

## install libosrm
```bash
git clone https://github.com/Project-OSRM/osrm-backend.git
mkdir -p build
cd build
cmake ..
cmake --build .
sudo cmake --build . --target install
```

## configure cmake 
make sure pkg-config works for libosrm 
```
echo $PKG_CONFIG_PATH
pkg-config libosrm --libs
pkg-config libosrm --cflags
```
if it doesn't work, check if libosrm.pc is included in PKG_CONFIG_PATH

