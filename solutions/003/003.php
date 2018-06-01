<?php

function largest_prime_factor($number) {
    $factor = 1;
    for ($i = 3; $i * $i <= $number; $i += 2) {
        if ($number == 1)
            return $factor;
        else if ($number % $i != 0)
            continue;
        else {
            $factor = $i;
            while ($number % $i == 0)
                $number /= $i;
        }
    }
    return $number;
}

echo largest_prime_factor(600851475143) . "\n";
