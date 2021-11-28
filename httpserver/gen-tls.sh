openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=cncamp.com/O=cncamp"
KEY=$(cat tls.key| base64 | tr -d "\n")
CRT=$(cat tls.crt| base64 | tr -d "\n")
sed -i "s/tls.crt: .*$/tls.crt: $CRT/g" secret.yaml
sed -i "s/tls.key: .*$/tls.key: $KEY/g" secret.yaml
