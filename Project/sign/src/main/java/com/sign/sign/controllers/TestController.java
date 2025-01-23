package com.sign.sign.controllers;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;



@Controller
public class TestController {
    @GetMapping("/index")
    public String index() {
        return "index";
    }

    @GetMapping("/generate_key")
    public String generateKey() {
        return "generate_key";
    }
    
    
}
