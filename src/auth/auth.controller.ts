import { Controller, Post, Body, HttpCode } from '@nestjs/common';
import { AuthService } from './auth.service';
import { RegisterRequestDto } from './dto/register-request.dto';
import { LoginRequestDto } from './dto/login-request.dto';
import { CommonResponse } from 'src/common/dto/common-response.dto';
import {
  ApiExtraModels,
  ApiOperation,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';
import { RegisterResponseDto } from './dto/register-response.dto';
import { LoginResponseDto } from './dto/login-response.dto';
import { swaggerExampleUrls } from 'src/config/swagger-example-urls';

@ApiTags('Authentication')
@ApiExtraModels(
  CommonResponse,
  RegisterRequestDto,
  RegisterResponseDto,
  LoginRequestDto,
  LoginResponseDto,
)
@Controller('auth')
export class AuthController {
  constructor(private readonly authService: AuthService) {}

  @Post('register')
  @HttpCode(201)
  @ApiOperation({ summary: 'Register a new user' })
  @ApiResponse({
    status: 201,
    description: 'User successfully registered',
    schema: {
      example: {
        code: 0,
        message: 'Success',
        data: {
          id: '33f01179-9d75-4062-9012-591b54a25f64',
          nickname: '우아한하늘빛양치기',
          email: 'johndoe@example.com',
          profileImage: swaggerExampleUrls.profileImage,
          role: 'user',
          createdAt: '2024-11-12T12:00:00Z',
          updatedAt: '2024-11-12T12:00:00Z',
        },
      },
    },
  })
  @ApiResponse({
    status: 400,
    description: 'Validation error',
    schema: {
      example: {
        code: 5,
        message: 'Validation failed',
        details: [
          {
            field: 'email',
            message: 'email must be an email',
          },
        ],
      },
    },
  })
  @ApiResponse({
    status: 409,
    description: 'User already exists',
    schema: {
      example: {
        code: 1,
        message: 'User already exists',
      },
    },
  })
  @ApiResponse({
    status: 500,
    description: 'Internal server error',
    schema: {
      example: {
        code: 6,
        message: 'Internal server error',
      },
    },
  })
  async register(
    @Body() registerDto: RegisterRequestDto,
  ): Promise<CommonResponse<RegisterResponseDto>> {
    return this.authService.register(registerDto);
  }

  @Post('login')
  @HttpCode(200)
  @ApiOperation({ summary: 'Login a user' })
  @ApiResponse({
    status: 200,
    description: 'User successfully logged in',
    schema: {
      example: {
        code: 0,
        message: 'Success',
        data: {
          tokenType: 'Bearer',
          expiresIn: 86400,
          accessToken: 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...',
        },
      },
    },
  })
  @ApiResponse({
    status: 401,
    description: 'Invalid credentials',
    schema: {
      example: {
        code: 2,
        message: 'Invalid credentials',
      },
    },
  })
  @ApiResponse({
    status: 500,
    description: 'Internal server error',
    schema: {
      example: {
        code: 6,
        message: 'Internal server error',
      },
    },
  })
  async login(
    @Body() loginDto: LoginRequestDto,
  ): Promise<CommonResponse<LoginResponseDto>> {
    return this.authService.login(loginDto);
  }
}
