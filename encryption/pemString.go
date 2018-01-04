package encryption

var PKCS1NoPassPrivateBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQCuNrKBiKXgACMd14h4Ep2HyBdIVTzRfeVOUMv4g/h9nLchyuAL
gYMfF24n6VKzS+l5lux1QFTtKbSxZv33X3mPoZNYjA7ohoqbruHaCFe/KjGEj5AY
SCfVjzyBK8DTN5hwze0TKpDPq1hUlXb655iDxexwOttBcZ9ftKvG4dF3ywIDAQAB
AoGBAItpcA3iqUMJORDfcvELjI6lG6ShuWq3wbGbTK7SBR+YMqDTNdlfSeuul1NP
DwQ2Ql8v3Iez9IJZdqTkk2XLPH3WgfkX4hPZBvLDmS9WWgPJ9a9HlQ0K41g9cAbq
UHVlGWFxfO9Bx2RFjPpK7anRfGmc5NkSUT+rUAxv98eLU1A5AkEA4b5v9IRYIdRW
tq7Yjv9gmg5x6IbiJ4CLifSO7dVJMa7DvzjQgB/Chs7GIGuJFU2dDFRZ27eqFYMi
eGAvcT28/wJBAMWQMQa9ze02uhHAzFagrOr4i7egPGZ7yA1VebsGrP1zIEHBf+/z
/fjPNCTOWLiXZUzy2vwhRc9b+cGcxmC7qTUCQQC7qwkBuyBVh2RLXRIV10Kk8EtM
Jw5ODrRyjMhh+kVuMz7ycrYLPdwHHFFj1GJTPoHIJTvKyGZymeEHN9dZCh1bAkEA
nNiS86Uwnck7A+kq0QOTBKD6W3gtO7e95Ugc8qykHjuOOmYi6h+tXCE0Zpwjt7kK
RYLw9AZ+Y9fDe0C6BgY2mQJAZ6N8Q8yPoFzKcrcgJQlL8naRFXtiLPsqbRMkO/oB
2BUr2vzxogYRl1+9KeTJmRB2bd8Rb4TSib9WMqHe7aL2wA==
-----END RSA PRIVATE KEY-----`)

var PKCS1NoPassPublicBytes = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCuNrKBiKXgACMd14h4Ep2HyBdI
VTzRfeVOUMv4g/h9nLchyuALgYMfF24n6VKzS+l5lux1QFTtKbSxZv33X3mPoZNY
jA7ohoqbruHaCFe/KjGEj5AYSCfVjzyBK8DTN5hwze0TKpDPq1hUlXb655iDxexw
OttBcZ9ftKvG4dF3ywIDAQAB
-----END PUBLIC KEY-----`)

var PKCS1WithPassPrivateBytes = []byte(`-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,C30102F8469F46DDF4FC96CECE61D464

4KrKHsik2MoE3heHXY3+aQ/Om9EdlgiSP434nan1iL+uJ+d83eflD3SJdFdUX0j3
KWe8PSBcTYKyfZDfyLKw1bQG+GLbfQEfl4Z4qmvbBCU8nOaSzSoaa2MpP49hK14C
vUXFNq2oYqe7HBj/UbkT9/fvW3Wf09mwBGfi2yHpoTD5Qn5iJiAsECWUDouOUlEn
0X6NkCV5+UzvNvvfp4zTv6PnE5l3j63NFM64mn3B5RlragHagLzB5zwdGNKvSNGW
kSSK/1y+Iy5rGIAgAUmk4+c4OqzmHjhDq35acsrFuD9hHxKops2JGARpqtKlGhIR
A7m2sARCBt0MPVTStHKTO+MfuR6GY/69BhUA1161GIRB4a1aCCblzrpJFAcmZbSe
qt5kfKyLEXFafhcpB57riIJGMAEg1r4ZRmvjcC+FcEu2Abekifv9mB44UvZQxEeD
uRjHfb9QR05mrDJvnbP7IYWhXJZlUbmBpM02JXqFOapM0cIdwZmSDkVO4ls5rfSN
Zb4OmWqmG/USwGWr0oJbzSUZANeGLWBsgYiL1G2L1dobNIuUCtUdkY8HO8e7jGfQ
/XHC8DPfctZ1zp+Pcl4/UaP47ywkCs58zPd6z3DyhTm2iV1idEqmbsy9UiX5KKGb
9lNli0S4EBSckh+IALI/EKAnO5UjtLUWzAF/mzVP46BAOl2n+YnGdQHzp0yot/Dk
3NjVBPNjZq+qIE+W5/rEHKuSjFiSWIO7g4xOwBwBBAU8REzLwX8vijAkCqKAZ4k3
kLP5nJ7GT3PYoOT5zxoXoCjbJRlHP5jdGrak1EW7Hz+a0XTNKDlcTzioG24Wdce1
-----END RSA PRIVATE KEY-----`)

var PKCS1WithPassPublicBytes = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDcWy9I21aXjZfZT7qzQKJojuwC
QkuKTVELxatOwTkqIniOe4I1cjr8e+1tBzAk8cU3flpVMxD43zLB5wpOwVB8q09R
lXmdAceHC3zFqcow/bm8gCSKnQq8WIYH19fH+iaNcuWVN9h0U4yCCAydAYA7MUJs
98quVeyU8jLS6rs/fwIDAQAB
-----END PUBLIC KEY-----`)

var PKCS8PrivateBytes = []byte(`-----BEGIN PRIVATE KEY-----
MIICeQIBADANBgkqhkiG9w0BAQEFAASCAmMwggJfAgEAAoGBANxbL0jbVpeNl9lP
urNAomiO7AJCS4pNUQvFq07BOSoieI57gjVyOvx77W0HMCTxxTd+WlUzEPjfMsHn
Ck7BUHyrT1GVeZ0Bx4cLfMWpyjD9ubyAJIqdCrxYhgfX18f6Jo1y5ZU32HRTjIII
DJ0BgDsxQmz3yq5V7JTyMtLquz9/AgMBAAECgYEAgpen1k3na7oGnEucIk/T8JOD
W3NewYBA3/EuLWZmMoprWEXqTTJ9stTfSRZRehOvBcxkTMoO7O+abUA346S9xYro
iztLLudj0dKvSp+GqrG5YtFkmjswIrY5N6o1K7uhHAwIUftveatHKcZEcj+ZYMGK
ufv+xWuVqw/2zBWzE0ECQQD7VC2uH7dYG2co+WyowC1Aav61e0/07n21GLJJy+G5
MuOqzhNZQDF+0X3A86XtOEw76fRWxN+WMR3Nhm0oQF09AkEA4HOje6LOmmpqgaRt
NQWEm2A3ruBpo8VcY3DUB44na/JmXCBg/pd9PJn5lzRvAujBRpQIKokKhz0wIfjT
/JfTawJBAOllkgpa38wNnq55e0P3O87PzfxBKM6fz9OmvKM1mEiBjtWnDGbc61oH
OA5A9j0nA/y71jHiIHT85d1gUeyIYU0CQQDKBAq+urhNjz4xVTbdpAvIdP9pLJ6o
Qnh3IBQWgMHGjLPIc3QZcWvM4aEdkJnh+nALAC2haxrIwi/SQ8046cXlAkEAm2Wg
3GaiS0EdqwgMpfbiW6aQJsmQwmovHQl2aAVb+hB6i6hoCILl/nzrguYxgz65RD/Y
mVqX2oex936FgKh06A==
-----END PRIVATE KEY-----`)

var PKCS8PublicBytes = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDcWy9I21aXjZfZT7qzQKJojuwC
QkuKTVELxatOwTkqIniOe4I1cjr8e+1tBzAk8cU3flpVMxD43zLB5wpOwVB8q09R
lXmdAceHC3zFqcow/bm8gCSKnQq8WIYH19fH+iaNcuWVN9h0U4yCCAydAYA7MUJs
98quVeyU8jLS6rs/fwIDAQAB
-----END PUBLIC KEY-----`)