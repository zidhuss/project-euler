<?php

$fibo = array(0, 1);
for ($i = 2; $i < 34; $i++)
    array_push($fibo, $fibo[$i-2] + $fibo[$i-1]);

echo(array_sum(array_filter(
    $fibo, function($x) { return $x % 2 == 0; })) . "\n");
