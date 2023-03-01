#version 300 es

in highp vec2 vertex_position;
in mediump vec3 vertex_color;

out mediump vec3 color;

void main() {
    color = vertex_color;
    gl_Position = vec4(vertex_position, 1.0, 1.0);
}