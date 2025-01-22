package com.sign.sign.controllers;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.bind.annotation.GetMapping;



@RestController
public class TestController {
    
    @GetMapping("/test")
    public String test() {
        RestTemplate restTemplate = new RestTemplate();
        String goApiUrl = "http://localhost:8080/api/test";
        String response = restTemplate.getForObject(goApiUrl, String.class);
        return response;
    }
    
}
