# 打包应用程序
fyne package -os linux -icon ./resource/logo.png

# 创建一个文件夹存放解压后的文件
mkdir kde_gemini

# 解压打包好后的程序
echo "开始解压 kde_gemini.tar.xz"
tar Jxvf kde_gemini.tar.xz -C ./kde_gemini --remove-files

# 将.desktop文件复制到解压好的文件夹下
cp -f ./kde_gemini.desktop ./kde_gemini/usr/local/share/applications

# 重新压缩
echo "重新压缩 kde_gemini.tar.xz"
tar Jcvf kde_gemini.tar.xz ./kde_gemini --remove-files

# 重新计算sha256sums
sha256sum kde_gemini.tar.xz
