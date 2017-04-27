data "mydns_a_record_set" "baidu" {
  host = "baidu.com"
}

output "baidu_addrs" {
  value = "${join("\n", data.mydns_a_record_set.baidu.addrs)}"
}
