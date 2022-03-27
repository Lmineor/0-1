参考： [AES五种加密模式](https://www.cnblogs.com/starwolf/p/3365834.html)

分组密码有五种工作体制

### 电码本模式（Electronic Codebook Book (ECB)）
这种模式是将整个明文分为若干相同的小段，然后对每一小段进行加密
### 密码分组链接模式（Cipher Block Chaining (CBC)）
这种模式是先将明文切分成若干小段，然后每一小段与初始块或者上一段的密文段进行异或运算后，再与密钥进行加密。  
### 计算器模式（Counter(CTR)）
### 密码反馈模式（Cipher FeedBack (CFB)）
### 输出反馈模式（Output FeedBack (OFB)）

