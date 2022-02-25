# Puzzle Challenge

## Run Unit Tests and Generate Code Coverage

```
make test
```
Code coverage of all the files can be seen by opening coverage.html in web browser.

## Start Using Puzzle

```
make run
```
This will start a REST router listening on port 8080

## Examples

```
// Get All the Puzzles
$curl localhost:8080/api/v1/puzzles          
["p1","p2","p3","p4","p5","p6"]                                                                                               

// Display contents of a specific puzzle
// TODO: Print the contents in a pretty format
$curl localhost:8080/api/v1/puzzles/p1
["a f a ","x o b ","o r c "]

$curl localhost:8080/api/v1/puzzles/p5
["k w h d e t a w s u w s i b o v z e z b ","l v d d c j z w o d g v e y s e k w w m ","o b w j e k y u m j m a f i s w o r u c ","q l h y x c p y e q p y t t a x u x k b ","e x v d h z v z n h q e c i h d l x c k ","m w t o e c m h g k b k x t e t h a p y ","j o y b e y i a y j n y w r q x v p o h ","e p r w m p b t z i c c x l g m f z i t ","s e i a m o q p v y x r s c r w t g m f ","z o z i y v e p y y w a i y p t m e s d ","p s c i p q w g l k d o w y s o b v j m ","t o u k i j h j q l l y z c u a l b q g ","x y k d r b x z v o j d s k x e q s p c ","p t e x s a m m l i c n z q j m y y k d ","s s n t w n i d c m x l f j g v m z y v ","e j e n m f w l s h h c m z u t x q j c ","x f s s h h a p x q y v l m h q n s s k ","y x q n i v v d k b b c e e y d r a n b ","l f i s g p p l m a z n m d z w k h s b ","n j q e g e m a j m i z x i o g x i x i "]

// Searching "string" in a puzzle
$curl localhost:8080/api/v1/puzzles/p5/net
{"Coordinates":{"Column":6,"Row":16},"Direction":"Horizontal Right","Number of Rotations":8}

$curl localhost:8080/api/v1/puzzles/p5/abc
"Word is not present in the puzzle!"
```
