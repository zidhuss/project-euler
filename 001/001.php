<?php

function sum($n, $k) {
	return intdiv($k * (intdiv($n-1, $k) * (intdiv($n-1, $k) + 1)), 2);
}

$solution = sum(1000, 3) + sum(1000, 5) - sum(1000, 15);
print("$solution\n");
