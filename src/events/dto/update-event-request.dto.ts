import { ApiPropertyOptional } from '@nestjs/swagger';
import { IsOptional, IsString, IsInt } from 'class-validator';
import { swaggerExampleUrls } from 'src/config/swagger-example-urls';

export class UpdateEventRequestDto {
  @ApiPropertyOptional({
    description: 'Updated title of the event',
    example: 'Updated Music Festival',
  })
  @IsOptional()
  @IsString()
  title?: string;

  @ApiPropertyOptional({
    description: 'Updated thumbnail URL for the event',
    example: swaggerExampleUrls.eventThumbnail,
  })
  @IsOptional()
  @IsString()
  thumbnail?: string;

  @ApiPropertyOptional({
    description: 'Updated location or venue of the event',
    example: 'Updated Central Park',
  })
  @IsOptional()
  @IsString()
  place?: string;

  @ApiPropertyOptional({
    description: 'Updated main cast or performers of the event',
    example: 'Updated Famous Band',
  })
  @IsOptional()
  @IsString()
  cast?: string;

  @ApiPropertyOptional({
    description: 'Updated age limit for the event (if any)',
    example: 21,
  })
  @IsOptional()
  @IsInt()
  ageLimit?: number;

  @ApiPropertyOptional({
    description: 'SVG data for the event design or layout',
    example:
      '<svg xmlns="http://www.w3.org/2000/svg" width="100" height="100"><circle cx="50" cy="50" r="40" stroke="black" stroke-width="3" fill="red" /></svg>',
  })
  @IsOptional()
  @IsString()
  svg?: string;

  @ApiPropertyOptional({
    description: 'The starting time for ticketing in ISO 8601 format',
    type: String,
    example: '2024-11-20T10:00:00.000Z',
  })
  @IsOptional()
  @IsString()
  ticketingStartTime?: Date;
}
