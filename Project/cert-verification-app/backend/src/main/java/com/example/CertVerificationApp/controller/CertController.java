package com.example.CertVerificationApp.controller;

import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/certificates")
public class CertController {

    @PostMapping("/verify")
    public String verifyCertificate(@RequestBody String certificate) {
        // Logic to verify the certificate goes here
        return "Certificate verification result";
    }
}