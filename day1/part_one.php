<?php

$moduleMasses = include __DIR__ . "/input.php";

$totalFuel = 0;

foreach ($moduleMasses as $moduleMass) {
    $totalFuel += floor($moduleMass / 3) - 2;
}

echo $totalFuel . "\n";
