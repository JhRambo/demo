<?php
// 客户端
class JsonRPC
{
    private $conn;
    public function __construct()
    {
        $this->conn = fsockopen('127.0.0.1', 8866, $errno, $errstr, 3);
        if (!$this->conn) {
            return false;
        }
    }
    public function Call($method, $params)
    {
        $err = fwrite($this->conn, json_encode(array(
            'method' => $method,
            'params' => array($params),
            'id'     => 666,
        )) . "\n");
        if ($err === false)
            return false;
        stream_set_timeout($this->conn, 0, 3000);
        $line = fgets($this->conn);
        if ($line === false) {
            return NULL;
        }
        return json_decode($line, true);
    }
}

$jsonrpc = new JsonRPC();
// $res = $jsonrpc->Call("user.SayHello", "我是php客户端");        //字符串
// $res = $jsonrpc->Call("user.GetUserData", array("Id" => 1));    //获取数据
$res = $jsonrpc->Call("user.AddUserData", array("Name"=>"张三"));    //添加数据
print_r($res);
