<?php
// router.php
if (preg_match('/\.(?:png|jpg|jpeg|gif)$/', $_SERVER["REQUEST_URI"])) {
    return false;    // serve the requested resource as-is.
} else {
    $say = getenv("SAY") ?: "Welcome";
    $name = getenv("NAME");
    echo "<p>${say} to ${name}</p>";
}
?>
