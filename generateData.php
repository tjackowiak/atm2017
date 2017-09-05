<?php

function send() {
  $data = (object)[
    "time" => mt_rand(15,60),
    "value" => mt_rand(-10,120),
    "used" => mt_rand(1,60),
  ];

  $options = [
    'http' => [
      'method'  => 'POST',
      'content' => json_encode( $data ),
      'header'=>  "Content-Type: application/json\r\n" .
                  "Accept: application/json\r\n"
      ]
  ];

  $url = "http://localhost:8080/result";

  $context  = stream_context_create( $options );
  $res = file_get_contents( $url, false, $context );

  if (!$res) {
    print_r($data);
  }
  return $res;
  // return json_decode( $result );
}

for ($i=0; $i < 200; $i++) {
  echo send(), PHP_EOL;
  usleep(mt_rand(10,100) * 10000);
}
