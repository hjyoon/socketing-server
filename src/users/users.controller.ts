import {
  Controller,
  Get,
  Body,
  Param,
  Delete,
  Patch,
  HttpCode,
  UseGuards,
} from '@nestjs/common';
import { UsersService } from './users.service';
import { CommonResponse } from 'src/common/dto/common-response.dto';
import { UpdateNicknameRequestDto } from './dto/update-nickname-request.dto';
import { UpdateNicknameResponseDto } from './dto/update-nickname-response.dto';
import { UpdatePasswordRequestDto } from './dto/update-password-request.dto';
import {
  ApiBearerAuth,
  ApiBody,
  ApiOperation,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';
import { JwtAuthGuard } from 'src/auth/jwt-auth.guard';
import { swaggerExampleUrls } from 'src/config/swagger-example-urls';

@ApiTags('Users')
@Controller('users')
export class UsersController {
  constructor(private readonly userService: UsersService) {}

  @ApiOperation({ summary: 'Retrieve user details' })
  @ApiResponse({
    status: 200,
    description: 'User details retrieved successfully',
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
    status: 404,
    description: 'User not found',
    schema: {
      example: {
        code: 7,
        message: 'User not found',
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
  @Get(':id')
  @HttpCode(200)
  findOne(@Param('id') id: string): Promise<CommonResponse<any>> {
    return this.userService.findOne(id);
  }

  @ApiOperation({ summary: 'Retrieve user details by email' })
  @ApiResponse({
    status: 200,
    description: 'User details retrieved successfully',
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
    status: 404,
    description: 'User not found',
    schema: {
      example: {
        code: 7,
        message: 'User not found',
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
  @Get('email/:email')
  @HttpCode(200)
  findOneByEmail(@Param('email') email: string): Promise<CommonResponse<any>> {
    return this.userService.findOneByEmail(email);
  }

  @ApiOperation({ summary: 'Update user nickname' })
  @ApiBearerAuth()
  @ApiBody({
    description: 'The new nickname for the user',
    type: UpdateNicknameRequestDto,
  })
  @ApiResponse({
    status: 200,
    description: 'Nickname updated successfully',
    schema: {
      example: {
        code: 0,
        message: 'Success',
        data: {
          id: '550e8400-e29b-41d4-a716-446655440000',
          nickname: 'new_nickname',
        },
      },
    },
  })
  @ApiResponse({
    status: 401,
    description: 'Unauthorized - Token is invalid or missing',
    schema: {
      example: {
        code: 8,
        message: 'Unauthorized',
      },
    },
  })
  @ApiResponse({
    status: 404,
    description: 'User not found',
    schema: {
      example: {
        code: 7,
        message: 'User not found',
      },
    },
  })
  @ApiResponse({
    status: 409,
    description: 'Nickname already exists',
    schema: {
      example: {
        code: 4,
        message: 'Nickname already exists',
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
  @Patch(':id/nickname')
  @HttpCode(200)
  @UseGuards(JwtAuthGuard)
  updateNickname(
    @Param('id') id: string,
    @Body() updateNicknameDto: UpdateNicknameRequestDto,
  ): Promise<CommonResponse<UpdateNicknameResponseDto>> {
    return this.userService.updateNickname(id, updateNicknameDto);
  }

  @ApiOperation({ summary: 'Update user password' })
  @ApiBearerAuth()
  @ApiBody({
    description: 'The new password for the user',
    type: UpdatePasswordRequestDto,
  })
  @ApiResponse({
    status: 200,
    description: 'Password updated successfully',
    schema: {
      example: {
        code: 0,
        message: 'Success',
      },
    },
  })
  @ApiResponse({
    status: 404,
    description: 'User not found',
    schema: {
      example: {
        code: 7,
        message: 'User not found',
      },
    },
  })
  @ApiResponse({
    status: 401,
    description: 'Unauthorized - Token is invalid or missing',
    schema: {
      example: {
        code: 8,
        message: 'Unauthorized',
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
  @Patch(':id/password')
  @HttpCode(200)
  @UseGuards(JwtAuthGuard)
  updatePassword(
    @Param('id') id: string,
    @Body() updatePasswordDto: UpdatePasswordRequestDto,
  ): Promise<CommonResponse<any>> {
    return this.userService.updatePassword(id, updatePasswordDto);
  }

  @ApiOperation({ summary: 'Soft delete a user' })
  @ApiBearerAuth()
  @ApiResponse({
    status: 204,
    description: 'User soft-deleted successfully',
  })
  @ApiResponse({
    status: 404,
    description: 'User not found',
    schema: {
      example: {
        code: 7,
        message: 'User not found',
      },
    },
  })
  @ApiResponse({
    status: 401,
    description: 'Unauthorized - Token is invalid or missing',
    schema: {
      example: {
        code: 8,
        message: 'Unauthorized',
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
  @Delete(':id')
  @HttpCode(204)
  @UseGuards(JwtAuthGuard)
  softDelete(@Param('id') id: string): Promise<any> {
    return this.userService.softDelete(id);
  }

  @ApiOperation({ summary: 'Retrieve user points by email' })
  @ApiResponse({
    status: 200,
    description: 'User points retrieved successfully',
    schema: {
      example: {
        code: 0,
        message: 'Success',
        data: {
          id: '33f01179-9d75-4062-9012-591b54a25f64',
          point: 100000,
        },
      },
    },
  })
  @ApiResponse({
    status: 404,
    description: 'User not found',
    schema: {
      example: {
        code: 7,
        message: 'User not found',
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
  @Get('/:id/points')
  @HttpCode(200)
  findUserPoints(@Param('id') id: string): Promise<CommonResponse<any>> {
    return this.userService.findUserPoints(id);
  }
}
