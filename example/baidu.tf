data "mydns_a_record_set" "baidu" {
  # resolver = "8.8.8.8:53"
  host = "baidu.com"
}

output "baidu_addrs" {
  value = "${join(":", data.mydns_a_record_set.baidu.addrs)}"
}

output "resolver" {
  value = "${data.mydns_a_record_set.baidu.resolver}"
}
