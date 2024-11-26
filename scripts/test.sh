#!/bin/sh

# Shell 环境

#!/bin/bash #! 告诉系统这个脚本需要什么解释器来执行，即使用哪一种 Shell。
echo "Hello World !" # echo 命令用于向窗口输出文本。
# bash ./test.sh #执行脚本
# sh ./test.sh #执行脚本





# Shell 变量

# 定义变量时，变量名不加美元符号（$），如：
your_name="qinjx"
echo $your_name
echo ${your_name} # 推荐

myUrl="https://www.google.com"
readonly myUrl # 只读变量
unset myUrl  # 删除变量

# 变量类型

# 字符串变量： 在 Shell中，变量通常被视为字符串。
my_string='Hello, World!'
my_string="Hello, World!"

# 整数变量： 在一些Shell中，你可以使用 declare 或 typeset 命令来声明整数变量。
declare -i my_integer=42 # 如果尝试将非整数值赋给它，Shell会尝试将其转换为整数。

# 数组变量： Shell 也支持数组，允许你在一个变量中存储多个值。
my_array=(1 2 3 4 5)

# 环境变量： 这些是由操作系统或用户设置的特殊变量，用于配置 Shell 的行为和影响其执行环境。
echo $PATH

# 特殊变量： 有一些特殊变量在 Shell 中具有特殊含义，例如 $0 表示脚本的名称，$1, $2, 等表示脚本的参数。
# $#表示传递给脚本的参数数量，$? 表示上一个命令的退出状态等。





# Shell 字符串

## 单引号
str='this is a string'
# 单引号里的任何字符都会原样输出，单引号字符串中的变量是无效的；
# 单引号字符串中不能出现单独一个的单引号（对单引号使用转义符后也不行），但可成对出现，作为字符串拼接使用。

## 双引号
your_name="runoob"
str="Hello, I know you are \"$your_name\"! \n"
echo -e $str
# 双引号里可以有变量
# 双引号里可以出现转义字符

## 拼接字符串
your_name="runoob"
# 使用双引号拼接
greeting="hello, "$your_name" !"
greeting_1="hello, ${your_name} !"
echo $greeting  $greeting_1 # hello, runoob ! hello, runoob !
# 使用单引号拼接
greeting_2='hello, '$your_name' !'
greeting_3='hello, ${your_name} !'
echo $greeting_2  $greeting_3 # hello, runoob ! hello, ${your_name} !

## 获取字符串长度
string="abcd"
echo ${#string}   # 输出 4 | 变量为字符串时，${#string} 等价于 ${#string[0]}:

## 提取子字符串
string="runoob is a great site"
echo ${string:1:4} # 输出 unoo | 从字符串第 2 个字符开始截取 4 个字符

## 查找子字符串
string="runoob is a great site"
echo `expr index "$string" io`  # 输出 4 | 查找字符 i 或 o 的位置(哪个字母先出现就计算哪个)





# Shell 数组

## 定义数组
array_name=(value0 value1 value2 value3)

## 读取数组
valuen=${array_name[n]}

## 获取数组的长度
# 取得数组元素的个数
length=${#array_name[@]}
length=${#array_name[*]}
# 取得数组单个元素的长度
length=${#array_name[n]}

## 关联数组
declare -A site=(["google"]="www.google.com" ["runoob"]="www.runoob.com" ["taobao"]="www.taobao.com")
echo ${site["runoob"]}





# Shell 传递参数
$#	## 传递到脚本的参数个数
$*	## 以一个单字符串显示所有向脚本传递的参数。如"$*"用「"」括起来的情况、以"$1 $2 … $n"的形式输出所有参数。
$$	## 脚本运行的当前进程ID号
$!	## 后台运行的最后一个进程的ID号
$@	## 与$*相同，但是使用时加引号，并在引号中返回每个参数。如"$@"用「"」括起来的情况、以"$1" "$2" … "$n" 的形式输出所有参数。
$-	## 显示Shell使用的当前选项，与set命令功能相同。
$?	## 显示最后命令的退出状态。0表示没有错误，其他任何值表明有错误。





# Shell 基本运算符

# expr 是一款表达式计算工具
val=`expr 2 + 2`
echo "两数之和为 : $val" # "两数之和为 : 4"

# 算术运算符
+	  # 加法	`expr $a + $b` 结果为 30。
-	  # 减法	`expr $a - $b` 结果为 -10。
*	  # 乘法	`expr $a \* $b` 结果为  200。
/	  # 除法	`expr $b / $a` 结果为 2。
%	  # 取余	`expr $b % $a` 结果为 0。
=	  # 赋值	a=$b 把变量 b 的值赋给 a。
==	# 相等。用于比较两个数字，相同则返回 true。	[ $a == $b ] 返回 false。
!=	# 不相等。用于比较两个数字，不相同则返回 true。	[ $a != $b ] 返回 true。

# 关系运算符
-eq	# 检测两个数是否相等，相等返回 true。	[ $a -eq $b ] 返回 false。
-ne	# 检测两个数是否不相等，不相等返回 true。	[ $a -ne $b ] 返回 true。
-gt	# 检测左边的数是否大于右边的，如果是，则返回 true。	[ $a -gt $b ] 返回 false。
-lt	# 检测左边的数是否小于右边的，如果是，则返回 true。	[ $a -lt $b ] 返回 true。
-ge	# 检测左边的数是否大于等于右边的，如果是，则返回 true。	[ $a -ge $b ] 返回 false。
-le	# 检测左边的数是否小于等于右边的，如果是，则返回 true。	[ $a -le $b ] 返回 true。

# 布尔运算符
!	  # 非运算，表达式为 true 则返回 false，否则返回 true。	[ ! false ] 返回 true。
-o	# 或运算，有一个表达式为 true 则返回 true。	[ $a -lt 20 -o $b -gt 100 ] 返回 true。
-a	# 与运算，两个表达式都为 true 才返回 true。	[ $a -lt 20 -a $b -gt 100 ] 返回 false。

# 逻辑运算符
&&	# 逻辑的 AND	[[ $a -lt 100 && $b -gt 100 ]] 返回 false
||	# 逻辑的 OR	[[ $a -lt 100 || $b -gt 100 ]] 返回 true

# 字符串运算符
=	  # 检测两个字符串是否相等，相等返回 true。	[ $a = $b ] 返回 false。
!=	# 检测两个字符串是否不相等，不相等返回 true。	[ $a != $b ] 返回 true。
-z	# 检测字符串长度是否为0，为0返回 true。	[ -z $a ] 返回 false。
-n	# 检测字符串长度是否不为 0，不为 0 返回 true。	[ -n "$a" ] 返回 true。
$	  # 检测字符串是否不为空，不为空返回 true。	[ $a ] 返回 true。

# 文件测试运算符
-b file	# 检测文件是否是块设备文件，如果是，则返回 true。	[ -b $file ] 返回 false。
-c file	# 检测文件是否是字符设备文件，如果是，则返回 true。	[ -c $file ] 返回 false。
-d file	# 检测文件是否是目录，如果是，则返回 true。	[ -d $file ] 返回 false。
-f file	# 检测文件是否是普通文件（既不是目录，也不是设备文件），如果是，则返回 true。	[ -f $file ] 返回 true。
-g file	# 检测文件是否设置了 SGID 位，如果是，则返回 true。	[ -g $file ] 返回 false。
-k file	# 检测文件是否设置了粘着位(Sticky Bit)，如果是，则返回 true。	[ -k $file ] 返回 false。
-p file	# 检测文件是否是有名管道，如果是，则返回 true。	[ -p $file ] 返回 false。
-u file	# 检测文件是否设置了 SUID 位，如果是，则返回 true。	[ -u $file ] 返回 false。
-r file	# 检测文件是否可读，如果是，则返回 true。	[ -r $file ] 返回 true。
-w file	# 检测文件是否可写，如果是，则返回 true。	[ -w $file ] 返回 true。
-x file	# 检测文件是否可执行，如果是，则返回 true。	[ -x $file ] 返回 true。
-s file	# 检测文件是否为空（文件大小是否大于0），不为空返回 true。	[ -s $file ] 返回 true。
-e file	# 检测文件（包括目录）是否存在，如果是，则返回 true。	[ -e $file ] 返回 true。





# Shell echo命令

# 1.显示普通字符串:
echo "It is a test" # echo It is a test

# 2.显示转义字符
echo "\"It is a test\"" # "It is a test"

# 3.显示变量
read name # read 命令从标准输入中读取一行,并把输入行的每个字段的值指定给 shell 变量
echo "$name It is a test"  # OK                     标准输入
                           # OK It is a test        输出

# 4.显示换行
echo -e "OK! \n" # -e 开启转义
echo "It is a test" #OK!
                    #
                    #It is a test

# 5.显示不换行
echo -e "OK! \c" # -e 开启转义 \c 不换行
echo "It is a test" # OK! It is a test

# 6.显示结果定向至文件
echo "It is a test" > myfile

# 7.原样输出字符串，不进行转义或取变量(用单引号)
echo '$name\"' # $name\"

# 8.显示命令执行结果
echo `date` # Thu Jul 24 10:08:46 CST 2014





# Shell printf 命令
# 默认的 printf 不会像 echo 自动添加换行符，我们可以手动添加 \n。

printf  format-string  [arguments...]
# format-string: 为格式控制字符串
# arguments: 为参数列表。

echo "Hello, Shell" # Hello, Shell
printf "Hello, Shell\n" # Hello, Shell

printf "%-10s %-8s %-4s\n" 姓名 性别 体重kg
printf "%-10s %-8s %-4.2f\n" 郭靖 男 66.1234
printf "%-10s %-8s %-4.2f\n" 杨过 男 48.6543
printf "%-10s %-8s %-4.2f\n" 郭芙 女 47.9876   #姓名     性别   体重kg
                                              #郭靖     男      66.12
                                              #杨过     男      48.65
                                              #郭芙     女      47.99
# %s %c %d %f 都是格式替代符，％s 输出一个字符串，％d 整型输出，％c 输出一个字符，％f 输出实数，以小数形式输出。

# %-10s 指一个宽度为 10 个字符（- 表示左对齐，没有则表示右对齐），任何字符都会被显示在 10 个字符宽的字符内，
# 如果不足则自动以空格填充，超过也会将内容全部显示出来。

# %-4.2f 指格式化为小数，其中 .2 指保留2位小数。

## printf 的转义序列
\a	  # 警告字符，通常为ASCII的BEL字符
\b	  # 后退
\c	  # 抑制（不显示）输出结果中任何结尾的换行字符，而且，任何留在参数里的字符、任何接下来的参数以及任何留在格式字符串中的字符，都被忽略
\f	  # 换页（formfeed）
\n	  # 换行
\r	  # 回车（Carriage return）
\t	  # 水平制表符
\v	  # 垂直制表符
\\	  # 一个字面上的反斜杠字符
\ddd	# 表示1到3位数八进制值的字符。仅在格式字符串中有效
\0ddd	# 表示1到3位的八进制值字符





# Shell test 命令
# Shell中的 test 命令用于检查某个条件是否成立，它可以进行数值、字符和文件三个方面的测试。

## 数值测试
-eq	# 等于则为真
-ne	# 不等于则为真
-gt	# 大于则为真
-ge	# 大于等于则为真
-lt	# 小于则为真
-le	# 小于等于则为真

num1=100
num2=100
if test $[num1] -eq $[num2]
then
    echo '两个数相等！'
else
    echo '两个数不相等！'
fi # 两个数相等！

#!/bin/bash
a=5
b=6
result=$[a+b] # 注意等号两边不能有空格
echo "result 为： $result" # result 为： 11

## 字符串测试
=	  # 等于则为真
!=	# 不相等则为真
-z  # 字符串	字符串的长度为零则为真
-n  # 字符串	字符串的长度不为零则为真

num1="ru1noob"
num2="runoob"
if test $num1 = $num2
then
    echo '两个字符串相等!'
else
    echo '两个字符串不相等!'
fi  # 两个字符串不相等!






# Shell 流程控制

# if
if condition
then
    command1
    command2
    ...
    commandN
fi

# if else
if condition
then
    command1
    command2
    ...
    commandN
else
    command
fi                    num1=$[2*3]
                      num2=$[1+5]
                      if test $[num1] -eq $[num2]
                      then
                          echo '两个数字相等!'
                      else
                          echo '两个数字不相等!'
                      fi

# for 循环
for var in item1 item2 ... itemN
do
    command1
    command2
    ...
    commandN
done                 for loop in 1 2 3 4 5
                     do
                         echo "The value is: $loop"
                     done

# while 语句
while condition
do
    command
done                int=1
                    while(( $int<=5 ))
                    do
                        echo $int
                        let "int++"
                    done

# 无限循环
while true
do
    command
done

# case ... esac
case 值 in
模式1)
    command1
    command2
    ...
    commandN
    ;;
模式2)
    command1
    command2
    ...
    commandN
    ;;
esac                        echo '输入 1 到 4 之间的数字:'
                            echo '你输入的数字为:'
                            read aNum
                            case $aNum in
                                1)  echo '你选择了 1'
                                ;;
                                2)  echo '你选择了 2'
                                ;;
                                3)  echo '你选择了 3'
                                ;;
                                4)  echo '你选择了 4'
                                ;;
                                *)  echo '你没有输入 1 到 4 之间的数字'
                                ;;
                            esac                                        输入 1 到 4 之间的数字:
                                                                        你输入的数字为:
                                                                        3
                                                                        你选择了 3

# 跳出循环
# break 命令 | break 命令允许跳出所有循环（终止执行后面的所有循环）。
while :
do
    echo -n "输入 1 到 5 之间的数字:"
    read aNum
    case $aNum in
        1|2|3|4|5) echo "你输入的数字为 $aNum!"
        ;;
        *) echo "你输入的数字不是 1 到 5 之间的! 游戏结束"
            break
        ;;
    esac
done                                    输入 1 到 5 之间的数字:3
                                        你输入的数字为 3!
                                        输入 1 到 5 之间的数字:7
                                        你输入的数字不是 1 到 5 之间的! 游戏结束

# continue | continue 命令与 break 命令类似，只有一点差别，它不会跳出所有循环，仅仅跳出当前循环。
while :  # 运行代码发现，当输入大于5的数字时，该例中的循环不会结束，语句 echo "游戏结束" 永远不会被执行。
do
    echo -n "输入 1 到 5 之间的数字: "
    read aNum
    case $aNum in
        1|2|3|4|5) echo "你输入的数字为 $aNum!"
        ;;
        *) echo "你输入的数字不是 1 到 5 之间的!"
            continue
            echo "游戏结束"
        ;;
    esac
done






# Shell 函数
funWithReturn(){
    echo "这个函数会对输入的两个数字进行相加运算..."
    echo "输入第一个数字: "
    read aNum
    echo "输入第二个数字: "
    read anotherNum
    echo "两个数字分别为 $aNum 和 $anotherNum !"
    return $(($aNum+$anotherNum))
}
funWithReturn
echo "输入的两个数字之和为 $? !"          这个函数会对输入的两个数字进行相加运算...
                                        输入第一个数字:
                                        1
                                        输入第二个数字:
                                        2
                                        两个数字分别为 1 和 2 !
                                        输入的两个数字之和为 3 !

# 函数参数
funWithParam(){
    echo "第一个参数为 $1 !"
    echo "第二个参数为 $2 !"
    echo "第十个参数为 $10 !"
    echo "第十个参数为 ${10} !"
    echo "第十一个参数为 ${11} !"
    echo "参数总数有 $# 个!"
    echo "作为一个字符串输出所有参数 $* !"
}
funWithParam 1 2 3 4 5 6 7 8 9 34 73      第一个参数为 1 !
                                          第二个参数为 2 !
                                          第十个参数为 10 !
                                          第十个参数为 34 !
                                          第十一个参数为 73 !
                                          参数总数有 11 个!
                                          作为一个字符串输出所有参数 1 2 3 4 5 6 7 8 9 34 73 !




# Shell 文件包含
. filename | source filename   # 注意点号(.)和文件名中间有一空格
# test1.sh : url="http://www.runoob.com"
# test1.sh : . ./test1.sh #

echo "菜鸟教程官网地址：$url"  # 菜鸟教程官网地址：http://www.runoob.com





# Shell 输入/输出重定向
# 一个命令通常从一个叫标准输入的地方读取输入，默认情况下，这恰好是你的终端。

command > file	# 将输出重定向到 file。
$ echo "菜鸟教程：www.runoob.com" > users
$ cat users
菜鸟教程：www.runoob.com

command >> file	# 将输出以追加的方式重定向到 file。
$ echo "菜鸟教程：www.runoob.com" >> users
$ cat users
菜鸟教程：www.runoob.com
菜鸟教程：www.runoob.com

command < file	  # 将输入重定向到 file。
$  wc -l < users  # 统计 users 文件的行数
       2

n > file	      # 将文件描述符为 n 的文件重定向到 file。
n >> file	      # 将文件描述符为 n 的文件以追加的方式重定向到 file。
n >& m	        # 将输出文件 m 和 n 合并。
n <& m	        # 将输入文件 m 和 n 合并。
<< tag	        # 将开始标记 tag 和结束标记 tag 之间的内容作为输入。