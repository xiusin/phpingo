<?php

$arr = [
    "xiusin",
    "xiusin",
    "yunjie",
    (object)[
        'name' => "xiusin"
    ],

    [
        'name' => "xiusin"
    ],
    [
        "age" => '23',
        'name' => 'yunjie'
    ]
];


// print_r(array_unique($arr));


print_r(array_splice($arr, 0, null, [1, 2, 3, 4]));;
print_r($arr);;
