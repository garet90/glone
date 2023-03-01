#version 300 es

in mediump vec3 color;

out mediump vec4 fragColor;

void main() {
    fragColor = vec4(color, 1.0);
}