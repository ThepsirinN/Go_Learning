module 0.3.2_call_module/hello

go 1.20

replace 0.3.1_create_module/greeting => ../0.3.1_create_module

require 0.3.1_create_module/greeting v0.0.0-00010101000000-000000000000
