#version 300 es

in highp vec2 vertex_position;

void main() {
    gl_Position = vec4(vertex_position, 1.0, 1.0);
}