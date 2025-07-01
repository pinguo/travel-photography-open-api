<?php

include "./SignatureBuilder.php";

$builder = new SignatureBuilder("ak","sk",300);
$rs = $builder->sign("/v1/api/test",["a"=>"B"],["c"=>"d"],"");
print_r($rs);
print_r($builder);