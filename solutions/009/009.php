<?php

function tripletProduct() {
    $s = 1000;
    $r = (int) sqrt($s);
    for ($m = $r; $m > 1; $m--)
        for ($n = $m; $n > 2; $n--) {
            $a = $m * $m - $n * $n;
            $b = 2 * ($m * $n);
            $c = $m * $m + $n * $n;
            if ($a + $b + $c === $s) {
                print($a * $b * $c . "\n");
                return;
            }
        }
}

tripletProduct();
