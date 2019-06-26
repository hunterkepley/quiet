package main

var (
	grayscaleShader = `
	#version 330 core
	in vec2  vTexCoords;
	out vec4 fragColor;
	uniform vec4 uTexBounds;
	uniform sampler2D uTexture;
	void main() {
		
		vec2 t = (vTexCoords - uTexBounds.xy) / uTexBounds.zw;
		
		float sum  = texture(uTexture, t).r;
			sum += texture(uTexture, t).g;
			sum += texture(uTexture, t).b;

		float gray = sum / 3.0f;
		
		vec4 color = vec4(gray, gray, gray, texture(uTexture, t).a);
		fragColor = color;
	}
	`

	stormShader = `
	#version 330 core
	in vec2  vTexCoords;
	out vec4 fragColor;
	uniform vec4 uTexBounds;
	uniform sampler2D uTexture;
	void main() {
		
		vec2 t = (vTexCoords - uTexBounds.xy) / uTexBounds.zw;
		
		vec4 color = vec4(texture(uTexture, t).r/3, texture(uTexture, t).g/3, texture(uTexture, t).b/2, texture(uTexture, t).a);
		fragColor = color;
		// For light v
		if((texture(uTexture, t).r == 0.99607843137 && texture(uTexture, t).g == 1.0 && texture(uTexture, t).b == 0.61568627451) 
		|| (texture(uTexture, t).r == 1.0 && texture(uTexture, t).g == 1.0 && texture(uTexture, t).b ==  0.78039215686)) {
			fragColor = vec4(texture(uTexture, t).r, texture(uTexture, t).g, texture(uTexture, t).b, texture(uTexture, t).a);
		}
	}
	`

	lightningShader = `
	#version 330 core
	in vec2  vTexCoords;
	out vec4 fragColor;
	uniform vec4 uTexBounds;
	uniform sampler2D uTexture;
	void main() {
		
		vec2 t = (vTexCoords - uTexBounds.xy) / uTexBounds.zw;
		
		vec4 color = vec4(texture(uTexture, t).r/3, texture(uTexture, t).g/3, texture(uTexture, t).b/2, texture(uTexture, t).a);
		fragColor = color;
		if(texture(uTexture, t).r >= 0.468 && texture(uTexture, t).g >= 0.507 && texture(uTexture, t).b >= 0.429) {
			fragColor = vec4(texture(uTexture, t).r, texture(uTexture, t).g, texture(uTexture, t).b, texture(uTexture, t).a);
		}
	}
	`

	regularShader = `
	#version 330 core
	
	in vec2  vTexCoords;
	out vec4 fragColor;
	uniform vec4 uTexBounds;
	uniform sampler2D uTexture;
	
	void main() {
		vec2 t = (vTexCoords - uTexBounds.xy) / uTexBounds.zw;
		fragColor = vec4(texture(uTexture, t).r, texture(uTexture, t).g, texture(uTexture, t).b, texture(uTexture, t).a);
	}
	`
)
