<?php
$dividend = 2;
$divisor = 0;
try{
	echo "$dividend 除以 $divisor ->\n";
	if($divisor == 0){
		throw new Exception('除数不能为零', 101);
	}
	$result = $dividend / $divisor;
	echo "结果为 $resultnn";
}
catch(Exception $e){
	# 类型指示 Exception
	$msg = $e->getMessage();
	$code = $e->getCode();
	$file = $e->getFile();
	$line = $e->getLine();
	$trace = $e->getTrace();
	// 注释
	$traceAsString = $e->getTraceAsString();
	echo "错误: $code, $msg\n";
	echo "错误所在文件: $file\n";
	echo "错误所在行号: $line\n";
	echo "回退路径数组: $trace\n";
	echo "回退路径字符串: $traceAsStringn\n";
}
