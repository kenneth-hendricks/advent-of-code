<?php

$moduleMasses = include __DIR__ . "/input.php";

$totalFuel = 0;

foreach ($moduleMasses as $moduleMass) {
    $totalFuel += getModuleFuel($moduleMass);
}

echo $totalFuel . "\n";


function getModuleFuel($mass)
{
    $fuel = floor($mass / 3) - 2;

    if ($fuel <= 0) {
        return 0;
    }

    return $fuel + getModuleFuel($fuel);
}
