data "mydns_a_record_set" "baidu" {
  host = "baidu.com"
}

output "baidu_addrs" {
  value = "${join(",", data.mydns_a_record_set.baidu.addrs)}"
}
