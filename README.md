glone is a wrapper for WebGL2 (OpenGL ES 2.0) in go. You can create a full rendering context by passing in the js.Value. See test/webgl for an example.

An application level de-duplicator is also included in the "dedup" package. Wrap any object that implements glone.RenderingContext to stop it from passing superfluous calls to the driver.